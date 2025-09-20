# Models
Models are structs that 'things' in our app. E.g., a chat message, a user, and so on.

Models are not directly stored in the DB. They can have related classes in the `infrastructure` package that will define how the data is store in the DB. This could be a generate class via `sqlc`.

