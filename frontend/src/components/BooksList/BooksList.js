import React from "react";
import { useHistory } from "react-router";
import ViewDetails from "./ViewDetails";
import axios from "axios";
import "./BooksList.css";
import BOOKSTOREIMAGE from "../../assets/bookstoreicon.jpeg";
import useAtomsApi from "../../hooks/useAtoms";

const BooksList = ({ booksList }) => {
  const history = useHistory();
  const { tokensData } = useAtomsApi();

  const handleBuy = (book) => {
    axios
      .post(
        "http://localhost:8080/api/auth/buy-book",
        {
          id: book.id,
          quantity: 1,
        },
        {
          headers: { denvar: tokensData.token },
        }
      )
      .then((response) => {
        history.push("/orderConfirmation", {
          book,
          isOrderPlaced: true,
        });
      })
      .catch(() => {
        history.push("/orderConfirmation", {
          book,
          isOrderPlaced: false,
        });
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
