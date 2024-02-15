import React from "react";
import { useHistory } from "react-router";
import "./OrderConfirmation.css";

const OrderConfirmation = () => {
  const history = useHistory();
  const { isOrderPlaced, book } = history.location.state;
  const { name, price } = book;

  const handleGoHome = () => {
    history.push("/home");
  };

  const SuccessOrder = () => {
    return (
      <>
        <h2 style={{ color: "green" }}>Order Placed Successfully</h2>
        <div className="order-details">
          <p>You have successfully purchased the following book:</p>
          <p>
            <strong>Book Name: </strong>
            {name}
          </p>
          <p>
            <strong>Price: </strong>&#8377;{price}
          </p>
        </div>
        <div data-testid="success-message" className="message">
          <p>Thank you for your purchase!</p>
          <p>We hope you enjoy reading.</p>
        </div>
      </>
    );
  };

  const FailureOrder = () => {
    return (
      <>
        <h2 style={{ color: "red" }}>Order Failed</h2>
        <div className="order-details">
          <p>You order for the following book has failed:</p>
          <p>
            <strong>Book Name: </strong>
            {name}
          </p>
          <p>
            <strong>Price: </strong>&#8377;{price}
          </p>
        </div>
        <div data-testid="failure-message" className="message">
          <p>Sorry for the inconvenience!</p>
        </div>
      </>
    );
  };

  const renderComponent = () => {
    if (isOrderPlaced) {
      return <SuccessOrder />;
    } else {
      return <FailureOrder />;
    }
  };

  return (
    <div className="order-container">
      <div
        className="order-confirmation-container"
        data-testid="order-confirmation"
      >
        {renderComponent()}
        <button
          data-testid="home-button"
          className="home-button"
          type="button"
          onClick={handleGoHome}
        >
          Back to Home
        </button>
      </div>
    </div>
  );
};

export default OrderConfirmation;
