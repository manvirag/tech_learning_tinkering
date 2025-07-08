# Distributed transaction

[https://www.youtube.com/watch?v=eltn4x788UM](https://www.youtube.com/watch?v=eltn4x788UM)

[https://www.youtube.com/watch?v=7FgU1D4EnpQ](https://www.youtube.com/watch?v=7FgU1D4EnpQ)

Problem statement:

we have multiple databases and we want to do transactions including both. Since both have their different transaction, how do we combine them?

Same problem to understand this.:

here store and delivery are the two different service and having their own database.

![Untitled](Distributed%20transaction%204ca6d956b01246b1b05b9580f6978669/Untitled.png)

![Untitled](Distributed%20transaction%204ca6d956b01246b1b05b9580f6978669/Untitled%201.png)

Some Solution:

1. Two-Phase Commit.

[https://drive.google.com/file/d/1FJa0DQPOxVJP2kXdyZDSttdMDI4Q2fUx/view](https://drive.google.com/file/d/1FJa0DQPOxVJP2kXdyZDSttdMDI4Q2fUx/view)

![Untitled](Distributed%20transaction%204ca6d956b01246b1b05b9580f6978669/Untitled%202.png)

![Untitled](Distributed%20transaction%204ca6d956b01246b1b05b9580f6978669/Untitled%203.png)

![Untitled](Distributed%20transaction%204ca6d956b01246b1b05b9580f6978669/Untitled%204.png)