Requirements:

Like Jira: Detail about app

(
The app contains multiple boards to signify different projects
Each board contains different lists to signify sub-project
Each list contain different cards signifying smaller tasks
Each card can be assigned to a user or may remain unassigned
)

Requirements:

1. User: Each user should have a userId, name, email.
2. Board: Each board should have a id, name, privacy (PUBLIC/PRIVATE), url, members, lists
3. List: Each list should have a id, name and cards
4. Card: Each card should have a id, name, description, assigned user
5. We should be able to create/delete boards, add/remove people from the members list and modify attributes.
   Deleting a board should delete all lists inside it.
6. We should be able to create/delete lists and modify attributes. Deleting a list should delete all cards inside it.
7. We should be able to create/delete cards, assign/unassign a member to the card and modify attributes
8. We should also be able to move cards across lists in the same board
9. Ability to show all boards, a single board, a single list and a single card
10. Default privacy should be public
11. Cards should be unassigned by default
12. Ids should be auto-generated for board/list/card
13. URLs should get created based on the id
