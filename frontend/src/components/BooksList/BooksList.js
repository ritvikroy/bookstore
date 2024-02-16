import React from "react";
import { useHistory } from "react-router";
import ViewDetails from "./ViewDetails";
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
          <div key={book.id} className="book-tile">
            <div>
              <img className="book-image" src={BOOKSTOREIMAGE} alt="book" />
            </div>
            <div className="book-details-container">
              <div className="book-details-container-name">{book.name}</div>
              <div className="book-details-container-price">
                Price : &#8377;{book.price}
              </div>
              <ViewDetails book={book} />
            </div>
            <div className="book-tile-button">
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
