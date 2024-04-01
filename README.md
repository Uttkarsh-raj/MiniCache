# MiniCache
<br>
MiniCache is an in-memory solution for storing key-value pairs and lists, supporting operations like pushing data from both ends of the list. One of its additional features is the ability to persist data, ensuring that even after restarting the server, the data remains intact. This is achieved by storing data on disk whenever a SET operation is performed. Upon server restart, the database checks for the presence of these files and retrieves the data if they exist, preventing any loss of information.
<br><br>


<!--TABLE OF CONTENTS-->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a> 
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a> 
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
  </details>

  There are many great README templates available on GitHub;however, I didn't find one that really suited my needs so I created this one. I wanted to create a README template which satisfies the project.
  
Here's why:

- A project that solves a problem and helps others.
- One shouldn't be doing the same task over and over like creating a README from scratch.
- You should implement DRY principles to the rest of your life :smile:

Of course, no one template will serve all projects since your needs may be different. So i'll be  adding more in the near future.You may also suggest changes by forking this repo and creating a pull request or opening issue. Thanks to all the people have contributed to expanding this template!

<!--About the Project-->
  
## About The Project

MiniCache is an in-memory solution for storing key-value pairs and lists, supporting operations like pushing data from both ends of the list. One of its key features is the ability to persist data, ensuring that even after restarting the server, the data remains intact. This is achieved by storing data on disk whenever a SET operation is performed. Upon server restart, the database checks for the presence of these files and retrieves the data if they exist, preventing any loss of information.
  

### Built With

This section should list any major frameworks/libraries used to bootstrap your project.Leave any add-ons/plugins for the acknowledgement section. Here are a few examples.
<br><br>

<img height="100px" src="https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg"/>



<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!--GETTING STARTED-->

## Getting Started

To get started with your Golang application, follow these steps:

1. **Install Golang**: Download and install Golang from the [official website](https://golang.org/dl/).

2. **Set Up Your Workspace**: Create a new directory for your project and set your `GOPATH` environment variable to point to this directory.

3. **Initialize Your Project**: Inside your project directory, run the following command to initialize a new Go module:

   ```
   go mod init github.com/your-username/project-name
   ```
   After installing Golang, you can start running your Go project.
4. **Run without Debugging**: In your terminal, navigate to the directory containing your main Go file (usually named `main.go`). Then, run the following command to build and execute your Go application:
   ```
   go run main.go
   ```
   This command will compile and execute your Go program without generating a binary file.



### Installation 

Below is an example of how you can instruct your audience on installing and setting up your app.This template doesn't rely on any external dependencies or services.

1.Clone the repo 

git clone https://github.com/Uttkarsh-raj/MiniCache


2.Install the packages 

go mod tidy


<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Routes

- **GET to test the server : '/hello'**
  * Response as
  ```
  {
  "message": "Hello There !!",
  "success": true
  }
  ```

- **POST to SET values in db : '/'**
  * Request
   ```
   {
    "command":"SET avg 12"
   }
  ```
  * Response as
  ```
  {
  "data":{
    "key": "avg",
    "value": "12",
  }
  "message": "",
  "success": true,
  }
  ```

- **POST to Get values of a db : '/'**
  * Request
  ```
  {"command":"GET avg"}
  ```
  * Response as
  ```
  {
  "data": {
    "key": "avg",
    "value": [
      "12"
  ]
  },
  "message": "",
  "success": true
  }
  ```
- **POST to push Left on List : '/'**
  * Request
   ```
   {
    "command":"LPUSH numbers 22"
   }
  ```
  * Response as
  ```
  {
  "data": {
    "key": "numbers",
    "value": {
      "List": [
        "22"
      ]
    }
  },
  "message": "",
  "success": true
  }
  ```

- **POST to push Right on List : '/'**
  * Request
  ```
  {
  "command":"RPUSH numbers 21"
  }
  ```
  * Response as
  ```
  {
  "data": {
    "key": "numbers",
    "value": {
      "List": [
        "22",
        "21"
      ]
    }
  },
  "message": "",
  "success": true
  }
  ```
- **POST to Get member of List : '/'**
  * Request
  ```
  {
  "command":"LMEMBERS numbers"
  }
  ```
  * Response as
  ```
  {
  "data": {
    "key": "numbers",
    "value": {
      "List": [
        "22",
        "21"
      ]
    }
  },
  "message": "",
  "success": true
  }
  ```
- **POST to get ALL elements in List : '/'**
  * Request
  ```
  {
  "command":"LMEMBERS *"
  }
  ```
  * Response as
  ```
  {
  "data": {
    "key": "*",
    "value": [
      {
        "key": "numbers",
        "value": {
          "List": [
            "22",
            "21"
          ]
        }
      },
    ]
  },
  "message": "",
  "success": true
  }
  ```

<!--USAGE EXAMPLES-->

## Usage
This project can be used in various scenarios where a lightweight, in-memory database with persistence features is required. Some potential use cases include:

1. **Caching**: Storing frequently accessed data in memory for fast access.
2. **Session Management**: Storing session data for web applications.
3. **Queue Management**: Implementing queues for message processing.
4. **Configuration Storage**: Storing and retrieving application configuration settings.
5. **Temporary Data Storage**: Holding temporary data that needs to be retained across server restarts.

By providing a simple and efficient in-memory database with persistence, the project can cater to these and other similar use cases, offering a reliable data storage solution.

<!-- ROADMAP -->

## Roadmap

- [x] Add Changelog
- [x] Add back to top links
- [x] Add Additional Templates w/ Examples
- [ ] Add "components" document to easily copy & paste sections of the readme
- [ ] Multi-language Support
  - [ ] Hindi
  - [ ] English

  
<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!--CONTRIBUTING-->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire ,and create.Any contributions you make are *greatly appreciated*.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->

## License


<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->

## Contact
Uttkarsh Raj - https://github.com/Uttkarsh-raj <br>

Project Link: https://github.com/Uttkarsh-raj/MiniCache

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->

## Acknowledgments

Use this space to list resources you find helpful and would like to give credit to. I've included a few of my favorites to kick things off!

- [Choose an Open Source License](https://choosealicense.com)
- [GitHub Emoji Cheat Sheet](https://www.webpagefx.com/tools/emoji-cheat-sheet)
- [Malven's Flexbox Cheatsheet](https://flexbox.malven.co/)
- [Malven's Grid Cheatsheet](https://grid.malven.co/)
- [Img Shields](https://shields.io)
- [GitHub Pages](https://pages.github.com)
- [Font Awesome](https://fontawesome.com)
- [React Icons](https://react-icons.github.io/react-icons/search)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
