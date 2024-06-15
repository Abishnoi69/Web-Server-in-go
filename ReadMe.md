## WebRouter

WebRouter is a simple web server written in Go. It serves static files from the static directory and handles two routes: /hello and /form.


### Getting Started
To run the server, you need to have Go installed on your machine. If you don't have it installed, you can download it from the official [Go website](https://golang.org/dl/).  Once you have Go installed, you can clone this repository and navigate to the project directory:

```bash
git clone https://github.com/Abishnoi69/WebRouter.git
cd WebRouter
```

Then, you can run the server using the following command:

```bash
go run main.go
```

The server will start running on port 8080. You can access it by navigating to http://localhost:8080 in your web browser.

### Routes

The server has two routes:
- `/hello`: This route returns a simple "Hello, World!" message.
  - Handled by the HelloHandler function in the modules package.
  
- `/form`: This route returns an HTML form that allows users to submit their name. When the form is submitted, the server responds with a message that includes the user's name.
    - Handled by the FormHandler function in the modules package.

## Static Files
The static files served by the server are located in the static directory. The home page (index.html) contains a welcome message and a link to a contact form.

## Dependencies
The server uses the following Go packages:
- `net/http`: Provides HTTP client and server implementations.
- `html/template`: Provides data-driven templates for generating HTML output.
- `log`: Provides a simple logging package.
- Many More .....

## Contributing
If you find a bug or have an idea for a new feature, feel free to open an issue or submit a pull request. We welcome contributions from the community and are happy to help you get started.
