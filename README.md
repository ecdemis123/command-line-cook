# command-line-cook
Want to figure out what to make for dinner, but don't feel like opening a browser? Now you can get tasty recipes right from the command line, like the real hacker you are.

## Background
command-line-cook gets its data from the [edamam api](https://developer.edamam.com/edamam-recipe-api).  You use command-line-cook like you would any other command line tool, just build and run it!

## Instructions
1.  Get API credentials from Edamam
2.  Store the credentials in your ~/.bashrc (or wherever you keep secrets)
```sh
export EDAMAM_APP_ID={app id}
export EDAMAM_APP_KEY={app key}
```
3. Run the executable ./command-line-cook. You can also copy into your /bin/ folder as well. command-line-cook takes one argument, the search param supplied as a flag.  This can be any sort of keyword for a recipe that you are looking for. Example:
```sh
$ command-line-cook -search=cheese
> Name: Skillet Mac And Cheese
Yield: 4 Calories: 2086
Ingredients:
* 2 cup macaroni
* 2 cup milk
* 3 tbsp all-purpose flour
* 1/4 tsp teaspoon paprika
* 2 cup grated cheddar cheese
* salt and white pepper to taste
* salt and white pepper to taste
Instructions:
http://www.sevenspoons.net/blog/2008/8/12/quick-fixes-reviewing-delias-how-to-cheat-at-cooking-and-eve.html
```
4. Get Cooking!

### Future Work
Update the tool to scrape the returned URL for recipe instructions
