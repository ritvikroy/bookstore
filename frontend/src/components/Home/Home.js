import React, { useEffect, useState } from "react";
import BooksList from "../BooksList/BooksList";
import SearchBar from "../SearchBar/SearchBar";
import axios from "axios";

const Home = () => {
  const [booksList, setBooksList] = useState([]);

  const getBooks = (text = "") => {
    const params = text ? `?searchText=${text}` : "";
    return axios.get("http://localhost:8080/api/auth/books" + params);
  };

  useEffect(() => {
    getBooks()
      .then((response) => {
        setBooksList(response?.data?.books);
      })
      .catch();
  }, []);

  const onSearch = (text) => {
    getBooks(text)
      .then((response) => {
        setBooksList(response?.data?.books);
      })
      .catch();
  };

  return (
    <div style={{ display: "flex", flexDirection: "column" }}>
      <SearchBar onSearch={onSearch} />
      <BooksList booksList={booksList} />
    </div>
  );
};

export default Home;
