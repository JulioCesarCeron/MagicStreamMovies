import { useState, useEffect } from "react"
import axiosConfig from "../../api/axiosConfig"
import Movies from "../movies/Movies.jsx"

const Home = () => {
  const [movies, setMovies] = useState([])
  const [loading, setLoading] = useState(false)
  const [message, setMessage] = useState("")

  useEffect(() => {
    const fetchMovies = async () => {
      setLoading(true)
      setMessage("")

      try {
        const response = await axiosConfig.get("/movies")
        setMovies(response.data)

        if (response.data.length === 0) {
          setMessage("There are currently no movies avaliable")
        }
      } catch (error) {
        console.log("Error fetching movies: ", error)
        setMessage("Error fetching movies")
      } finally {
        setLoading(false)
      }
    }

    fetchMovies()
  }, [])

  return (
    <>
      {loading ? (
        <h2>Loading...</h2>
      ) : (
        <Movies movies={movies} message={message} />
      )}
    </>
  )
}

export default Home
