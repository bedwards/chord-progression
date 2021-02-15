package main

import (
    "flag"
    "fmt"
    "os"

    "jalopymusic.com/jalopy/numerals"
    "jalopymusic.com/jalopy/tonality"
)

func main() {
    songCmd := flag.NewFlagSet("song", flag.ExitOnError)
    songKey := songCmd.String("key", "D#", "key")
    songParts := songCmd.String("parts", "aba", "parts")


    if len(os.Args) < 2 {
        fmt.Println("expected 'song' subcommand")
        os.Exit(1)
    }

    switch os.Args[1] {

    case "song":
        songCmd.Parse(os.Args[2:])
        if *songParts != "aba" {
            fmt.Printf("Unsupported song definition: %#v\n", songParts)
            os.Exit(1)
        }
        n := numerals.New()
        song, err := n.CreateSongABA()
        if err != nil {
            fmt.Printf("Song is broken: %#v\n", err)
            os.Exit(1)
        }

        fmt.Printf("%#v\n", song.A[:4])
        fmt.Printf("%#v\n", song.A[4:8])
        fmt.Printf("%#v\n", song.A[8:12])
        fmt.Printf("%#v\n", song.A[12:16])
        fmt.Printf("----\n")
        fmt.Printf("%#v\n", song.B[:4])
        fmt.Printf("%#v\n", song.B[4:8])
        fmt.Printf("%#v\n", song.B[8:12])
        fmt.Printf("%#v\n", song.B[12:16])
        fmt.Printf("----\n")
        
        t := tonality.New()
        a, err := t.InKey(*songKey, song.A)
        if err != nil {
            fmt.Printf("Song is broken: %#v\n", err)
            os.Exit(1)
        }
        b, err := t.InKey(*songKey, song.B)
        if err != nil {
            fmt.Printf("Song is broken: %#v\n", err)
            os.Exit(1)
        }

        fmt.Println(a)
        fmt.Println("----")
        fmt.Println(b)

    default:
        fmt.Println("expected 'song' subcommand")
        os.Exit(1)
    }
}
