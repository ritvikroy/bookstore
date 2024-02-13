import "./App.css";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import SignIn from "./components/Auth/SignIn";

function App() {
  return (
    <div className="App" data-testid="home">
      <BrowserRouter>
        <Switch>
          <Route path="/" component={SignIn} />
        </Switch>
      </BrowserRouter>
    </div>
  );
}

export default App;
