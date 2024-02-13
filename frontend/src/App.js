import React from 'react';
import {BrowserRouter,Route, Switch} from "react-router-dom";
import SignIn from './components/Auth/SignIn'
import BooksList from './components//BooksList/BooksList';
import "./App.css";


function App() {
  return (
    <BrowserRouter>
    <React.StrictMode>
      <div className="App" data-testid="home">
        <Switch>
          <Route path="/" exact component={SignIn}/>
          <Route path="/books" component={BooksList}/>
        </Switch>
      </div>
         </React.StrictMode>
         </BrowserRouter>
  );
}

export default App;
