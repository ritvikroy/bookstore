import { fireEvent, render, screen } from "@testing-library/react";
import OrderConfirmation from "./OrderConfirmation";

const mockHistoryPush = jest.fn();
const mockHistory = {
  push: mockHistoryPush,
  location: {state: {book: {id: 1, name: "Book 1", price: 200}, isOrderPlaced: true}}
};
jest.mock("react-router", () => {
  const originalModule = jest.requireActual("react-router");
  return {
    ...originalModule,
    useHistory: () => mockHistory,
  };
});

describe("Order Confirmation", ()=>{
    
    it("should render order confirmation screen with success message",()=>{
        render(<OrderConfirmation/>);
        const linkElement = screen.getByTestId("success-message");
        expect(linkElement).toBeInTheDocument();
    });

    it("should render order confirmation screen with failure message",()=>{
      mockHistory.location.state.isOrderPlaced = false;
      render(<OrderConfirmation/>);
      const linkElement = screen.getByTestId("failure-message");
      expect(linkElement).toBeInTheDocument();
  });

    it("should redirect to home page when back to home button is clicked",()=>{
        render(<OrderConfirmation/>);
        const linkElement = screen.getByTestId("home-button");
        fireEvent.click(linkElement);
        expect(mockHistoryPush).toHaveBeenCalledTimes(1);
        expect(mockHistoryPush).toHaveBeenCalledWith("/books");
    });
});