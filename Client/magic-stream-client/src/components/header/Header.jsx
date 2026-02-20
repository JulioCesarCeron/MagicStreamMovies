import { useState } from "react"
import Button from "react-bootstrap/Button"
import Container from "react-bootstrap/Container"
import Nav from "react-bootstrap/Nav"
import Navbar from "react-bootstrap/Navbar"
import { NavLink, useNavigate } from "react-router-dom"

const Header = () => {
  const [auth, setAuth] = useState(false)
  const navigate = useNavigate()

  return (
    <Navbar
      bg="dark"
      variant="dark"
      expand="lg"
      stick="top"
      className="shadow-sm"
    >
      <Container>
        <Navbar.Brand>Magic Stream</Navbar.Brand>

        <Navbar.Toggle aria-controls="main-navbar-nav" />
        <Navbar.Collapse>
          <Nav className="me-auto">
            <Nav.Link as={NavLink} to="/">
              Home
            </Nav.Link>
            <Nav.Link as={NavLink} to="/recommended">
              Recommended
            </Nav.Link>
          </Nav>

          <Nav className="ms-auto align-items-center">
            {auth ? (
              <>
                <span>
                  Hello, <strong>Name</strong>
                </span>
                <Button variant="outline-light" size="sm">
                  Logout
                </Button>
              </>
            ) : (
              <>
                <Button
                  variant="outline-info"
                  size="sm"
                  className="me-2"
                  onClick={() => navigate("/login")}
                >
                  Login
                </Button>
                <Button
                  variant="outline-info"
                  size="sm"
                  className="me-2"
                  onClick={() => navigate("/register")}
                >
                  register
                </Button>
              </>
            )}
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  )
}

export default Header
