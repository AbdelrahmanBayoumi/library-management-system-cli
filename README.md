<h1 align='center'>  Library Management System CLI </h1>

Library management system using Golang


## Requirements
* [go 1.15.3](https://golang.org/)



## Getting Started

1- Start the server.go, that is located in sever dirictory 
    
    go run server.go

3- You can now use the client cli [(how to use the client)](/README.md#books-operations-)
    
    go run server.go --help




## Books Operations 📚

#### 1- add book (returns the added book)

    go run client.go --book --add '{"id":"1","title":"title_01","publicationDate":"publicationDate_01","author":"author_01","genre":"genre_01","publisher":"publisher_01","language":"language_01"}'
    
#### 2- search book by id (returns the book if found)

    go run client.go --book --search --id "{'id' : '3'}"
    
#### 3- search book by title (returns the book if found)

    go run client.go --book --search --title "{'title' : 'Gardening Advice '}"

#### 4- get all books (returns list of books)

    go run client.go --books
    
#### 5- sort by title (returns list of books sorted by title)

    go run client.go --books --sort --title
    

#### 6- sort by publication_date (returns list of books sorted by publication_date)

    go run client.go --books --sort --publication_date
    
<hr>

## Readers Operations 👴🏻📖

#### 1- add a reader (returns the added reader)

    go run client.go --reader --add "{ "readers": [ { "id": "1", "name": "name_01", "gender": "gender_01", "birthday": "birthday_01", "height": "height_01", "weight": "weight_01", "employment": "employment_01" } ]}"
    
#### 2- remove reader by id (returns the removed reader if found)

    go run client.go --reader --remove --id "{'id' : '3'}"
    
#### 3- search reader by id (returns the reader if found)

    go run client.go --reader --search --id "{'id' : '3'}"
    
#### 4- search book by name (returns the reader if found)

    go run client.go --reader --search --title "{'title' : 'Gardening Advice '}"

#### 5- get all readers (returns list of readers)

    go run client.go --readers
    
  
  
  
## Contributing 💡
If you want to contribute to this project and make it better with new ideas, your pull request is very welcomed.
If you find any issue just put it in the repository issue section, thank you.


## Thank You!
Please ⭐️ this repo and share it with others


<div align='center'>
    <img alt="golang-logo" src="https://user-images.githubusercontent.com/48678280/103093126-20da7a00-4602-11eb-88ab-0903f976509b.png">
</div>



<br>

-----------

<h6 align="center">سبحَانَكَ اللَّهُمَّ وَبِحَمْدِكَ، أَشْهَدُ أَنْ لا إِلهَ إِلأَ انْتَ أَسْتَغْفِرُكَ وَأَتْوبُ إِلَيْكَ</h6>
