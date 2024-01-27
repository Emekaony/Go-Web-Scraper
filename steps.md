## Steps I took to build this projects and ideas:

1. I used the colly library to parse thru the html file and get the author name and date.
2. Now we need to make this whole thing extensible 
   1. maybe make a model for each thing we scrape
   2. like a struct with fields matching data we are interested in, and maybe interfaces and such
3. Using the encoding/csv package, we were able to save the blog data into a csv file. 
4. I also built custom functionality to properly parse web scraped data in a more readable and understandable format.