# Udemy Hangman Homework Assignment

## Installing dependencies on OS X

* ./install-dependencies

## Testing

* ./test will run basic testing infrastructure

## Environment Variables

* export WORD_FILE_PATHNAME=/usr/share/dict/words
  * Unix dict file on Mac OS X

## Unauthenticated Endpoints

* POST /games
  * create initial game state and redirect to new game on success

  * headers: nothing special
  
  * internally calls:
    * PUT /games/:game-id
    * PUT /games/:game-id/secret-word

  * possible responses
    * 302, Location: /games/:game-id
    * 500

* GET /games/:game-id
  * get current game state

  * possible responses
    * 200, body: Game entity
    * 404, game not found
    * 500

* POST /games/:game-id
  * make a guess

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

## Authenticated Endpoints

* PUT /games/:game-id
  * create or update game state, store history via ETag

  * headers:
    * If-Match:, for optimistic locking via ETag

  * body: Game entity

  * possible responses
    * 204, body: empty
    * 406, validation failed
    * 409, optimistic locking failed
    * 500

* PUT /games/:game-id/secret-word
  * create game's secret word

  * body: Game-Secret-Word entity

  * possible responses
    * 204, body: empty
    * 406, validation failed
    * 409, word already set
    * 500

* GET /hames/:game-id/secret-word
  * 200, body: Game-Secret-Word entity
  * 404, game or game word not found
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
