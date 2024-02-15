import React from "react";
import { useHistory } from "react-router";
import "./BooksList.css";
import BOOKSTOREIMAGE from "../../assets/bookstoreicon.jpeg";

const BooksList = ({ booksList }) => {
  const history = useHistory();

  const handleBuy = (book) => {
    history.push("/orderConfirmation", {
      book,
      isOrderPlaced: Math.random() > 0.5 ? true : false,
    });
  };

  return (
    <div className="book-list-container" data-testid="books-list">
      <div>
        {booksList.map((book) => (
          <div
            key={book.id}
            style={{
              display: "flex",
              backgroundColor: "aliceblue",
              justifyContent: "space-around",
              margin: "20px",
              boxShadow: "0 4px 8px 0 grey, 0 6px 20px 0 grey",
              borderRadius: "20px",
            }}
          >
            <div style={{ padding: "10px" }}>
              <img
                style={{ width: "200px", height: "200px" }}
                src={BOOKSTOREIMAGE}
                alt="book"
              />
            </div>
            <div
              style={{
                width: "50%",
                display: "flex",
                flexDirection: "column",
                alignItems: "flex-start",
                justifyContent: "space-around",
              }}
            >
              <div style={{ fontSize: "28px", fontWeight: "bold" }}>
                {book.name}
              </div>
              <div style={{ fontSize: "22px" }}>
                Price : &#8377;{book.price}
              </div>
            </div>
            <div
              style={{
                width: "10%",
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
              }}
            >
              <button
                data-testid={`${book.id}`}
                type="button"
                onClick={() => handleBuy(book)}
              >
                Buy
              </button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default BooksList;
