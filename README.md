\# GitUndo



GitUndo is a simple CLI tool built in Go to simplify common Git undo operations.  

It provides easy commands to revert commits, unstage files, discard file changes, and more.



---



\## 🚀 Features

\- Undo the \*\*last commit\*\*

\- Undo \*\*changes in a file\*\*

\- \*\*Unstage\*\* a file from staging area

\- Revert a commit by hash

\- View recent commit logs



---



\## 🛠️ Installation \& Usage

Clone the repository and run with Go:



```bash

git clone https://github.com/vamsikrishnach25-star/gitundo.git

cd gitundo

go run main.go --help


gitundo last-commit



gitundo file <filename>


gitundo unstage <filename>



gitundo revert <commit-hash>



gitundo log





Contributions are welcome!

Feel free to fork the repo, open issues, and submit pull requests.


