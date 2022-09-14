package cube

func CubeVolume(x int) int {
	return x * x * x // which refers to x(cubed) that's how to calculate the cubic volume
}

// :=
// pwd : shows the modules directory
// ex.   /home/ibrahim/Desktop/go_math
// cd .. : inorder to go to the directory
// using the commaned <git init> will track the pointed folder inorder to have connection with the repo
// as well as this command just created a new hidden directory in our directory called <.git> as shown
// To see hidden directories we use the command <ls -a>
// Inorder to add the remote repo and give it a name we use the command :-
// git remote add origin https://github.com/Ibrahimabdelgawi/go_math.git
// .git is mandatory as well as the protocol which is the https
// <git remote -v> := to check the name and the url of the remote repository
// The next step is to move all these files to what's known as the (staging area)
// staging area : refers to the area where git starts tracking and saving changes that occurs in files
// By using the command : git add -A
// so that we moved all the files from the current directory to the staging area
// nxt is to commit(save) these changes
// but first we need to setup some configuration inorder for get to not to submit an error
// this is done by submitting the username and the email
// git config user.name "ibrahim"
// git config user.email "ibrahimabdelgawi@icloud.com"
