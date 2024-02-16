import React from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import SignIn from "./components/Auth/SignIn";
import Home from "./components/Home/Home";
import { Provider } from "jotai";
import OrderConfirmation from "./components/OrderConfirmation/OrderConfirmation";
import "./App.css";

function App() {
  return (
    <Provider>
      <BrowserRouter>
        <React.StrictMode>
          <div className="App" data-testid="home">
            <Switch>
              <Route path="/" exact component={SignIn} />
              <Route path="/home" component={Home} />
              <Route path="/OrderConfirmation" component={OrderConfirmation} />
            </Switch>
          </div>
        </React.StrictMode>
      </BrowserRouter>
    </Provider>
  );
}

export default App;
