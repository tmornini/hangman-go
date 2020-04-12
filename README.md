# Udemy Hangman Homework Assignment

## Installing dependencies on OS X

* ./install-dependencies

## Testing

* ./test will run basic testing infrastructure

## Environment Variables

* export WORD_FILE_PATHNAME=/usr/share/dict/words
  * Unix dict file on Mac OS X

## Endpoints

* GET /games/:game-id/secret-word
  * get-game-secret-word

  * headers:
    * Authorization:, token

  * 200, body: Game-Secret-Word entity
  * 404, game or game word not found
  * 500

* GET /games/:game-id
  * get-game

  * possible responses
    * 200, body: Game entity
    * 404, game not found
    * 500

* POST /games/:game-id
  * post-game (guess a letter)

  * headers:
    * If-Match:, for optimistic locking via ETag, passes through to PUT /games/:game-id

  * body: Game-Guessed-Letter entity

  * internally calls:
    * GET /games/:game-id
    * GET /games/:game-id/secret-word
    * PUT /games/:game-id

  * possible responses
    * 302, Location: /games/:game-id
    * 400, deserialization failed
    * 404, game not found
    * 406, validation failed
    * 409, optimistic locking failed
    * 500

* POST /games
  * post-games (create a new game)

  * headers: nothing special
  
  * internally calls:
    * PUT /games/:game-id
    * PUT /games/:game-id/secret-word

  * possible responses
    * 302, Location: /games/:game-id
    * 500

* GET /games/:game-id/history
  * get-game-history

  * body: list of game states, oldest to newest

* PUT /games/:game-id
  * put-game

  * headers:
    * If-Match:, for optimistic locking via ETag
    * Authorization:, token

  * body: Game entity

  * possible responses
    * 204, body: empty
    * 406, validation failed
    * 409, optimistic locking failed
    * 500

* PUT /games/:game-id/secret-word
  * create game's secret word

  * headers:
    * Authorization:, token
    * If-None-Match:

  * body: Game-Secret-Word entity

  * possible responses
    * 204, body: empty
    * 406, validation failed
    * 409, word already set
    * 500

## Entity Specification

### Game

``` JSON
{
    "unmasked-letters": ["","a","","","","a",""],
    "wrong-guesses-remaining": 8,
    "guessed-letter-history": [
      { "guessed-letter": "a", "unmasked-count": 2 },
      { "guessed-letter": "b", "unmasked-count": 0 },
      { "guessed-letter": "c", "unmasked-count": 0 }
    ]
}
```

### Game-Secret-Word

``` JSON
{
    "secret-word": "hangman"
}
```

### Game-Guessed-Letter

``` JSON
{
    "guessed-letter": "c"
}
```
