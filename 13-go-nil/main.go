package main;

import "fmt";

func createNewMap(name string) map[string]string {
	if name == "" {
		return nil;
	}
	return map[string]string{"name": name};
}

func main() {
	fmt.Println("Create map with valid name");
	user1 := createNewMap("Umar");
	fmt.Println("User 1:", user1["name"]);
	fmt.Println("======================");

	fmt.Println("Create map with empty name");
	user2 := createNewMap("");
	if user2 == nil {
		fmt.Println("User 2 is empty");
	} else {
		fmt.Println("User 2:", user2["name"]);
	}
}