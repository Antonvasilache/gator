package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Antonvasilache/gator/internal/config"
	"github.com/Antonvasilache/gator/internal/database"
	_ "github.com/lib/pq"
)

func main(){
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}	

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to the db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)


	programState := &state{
		db: dbQueries,
		cfg: &cfg,
	}

	cmds := &commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("register", handlerRegister)
	cmds.register("login", handlerLogin)
	cmds.register("reset", handlerReset)
	cmds.register("users", handleGetUsers)
	cmds.register("agg", handleAgg)
	cmds.register("addfeed", middlewareLoggedIn(handleAddFeed))
	cmds.register("feeds", handleGetFeeds)
	cmds.register("follow", middlewareLoggedIn(handleFollow))
	cmds.register("following", middlewareLoggedIn(handleListFeedFollows))
	cmds.register("unfollow", middlewareLoggedIn(handleUnfollow))

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")		
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}