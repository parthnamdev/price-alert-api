# price-alert-api

[![Open in Visual Studio Code](https://classroom.github.com/assets/open-in-vscode-c66648af7eb3fe8bc4f294546bfd86ef473780cde1dea487d3c4ff354943c9ae.svg)](https://classroom.github.com/online_ide?assignment_repo_id=7943363&assignment_repo_type=AssignmentRepo)
<div id="top"></div>


<!-- PROJECT LOGO -->

<h3 align="center">Price alert API</h3>

  <p align="center">
    A CLI assistant to get alert on email when Bitcoin reaches certain price
    <br />

  </p>
</div>





<!-- ABOUT THE PROJECT -->
## About The Project


### Built With

* [GoLang](https://go.dev/)


<p align="right">(<a href="#top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

Make sure you have Go and MongoDB (or add Altas URI) installed in your system.

### Prerequisites

You can download Go if you don't have on your system
[Download GoLang](https://go.dev/dl/)

### Installation

1. Download Go
2. Clone the repo
   ```sh
   git clone https://github.com/parthnamdev/price-alert-api.git
   ```
3. Run Go command
   ```sh
   go run main.go
   ```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

1. /user/create POST
   ```sh
   {
    "username": "parth",
    "email": "parth@gmail.com",
    "password": "123456"
    }
    
2. /user/login POST
   ```sh
   {
    "username": "parth",
    "password": "123456"
  }
  
3. /user/home GET
  
4. /alert/create POST
   ```sh
   {
    "price": 25000,
    }
  
5. /alert/delete POST
   
_For more examples, please refer to the [Documentation]()_

<p align="right">(<a href="#top">back to top</a>)</p>







<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Your Name - [@_parthnamdev_](https://twitter.com/_parthnamdev_) [LinkedIn](https://linkedin.com/in/parth-namdev-4584331a2) - parthnamdevpm12345@gmail.com

Project Link: [https://github.com/parthnamdev/price-alert-api](https://github.com/parthnamdev/price-alert-api)

<p align="right">(<a href="#top">back to top</a>)</p>




<p align="right">(<a href="#top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/parth-namdev-4584331a2
