# Project Name

# Short Link

**Environment:**

- **Backend:** Built with Beego Framework (Golang).
    - Source code: [https://github.com/hvdtam/demo_shortLink_backend](https://github.com/hvdtam/demo_shortLink_backend)
- **Frontend:** Developed with NextJS.
    - **Language**: Typescript.
        - **CSS Framework**: TailwindCSS.
        - **State Management**: localStorage if not authentication. If authentication, using Server.
    - Source code: [https://github.com/hvdtam/demo_shortLink_frontend](https://github.com/hvdtam/demo_shortLink_frontend)
- **Database:** PostgresSQL.

---

## **Demo URL:**
Link: [https://s.tamk.dev](https://s.tamk.dev/)

**Function:**

- **Generate URL ShortLink:**
    - Enter custom alias for the shortlink
    - Set expiration time for the shortlink
    - Enable password protection for the shortlink
- **View:**
    - Copy shortlink to clipboard
    - Share shortlink with QR Code

**Shortlink Management:**

Authentication is required to manage ShortLinks.

**Authentication:**

Authentication using Bear Token.

## Project Structure

The project is structured as follows:

- `conf`: configuration files for the project.
- `controllers`: controllers for handling HTTP requests.
- `database`: database configuration and schema files.
- `helper`: utility functions and helper classes.
- `models`: data models for the project.
- `routers`: routers for mapping HTTP requests to controller functions.
- `static`: static files (images, CSS, JavaScript, etc.).
- `swagger`: Swagger API documentation files.
- `tests`: unit tests and integration tests for the project.
- `views`: HTML templates for rendering dynamic content.

## Installation

Config app.conf

## Usage

```bash
bee run
```
