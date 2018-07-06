import React, { Component } from 'react';
import { InstantSearch } from 'react-instantsearch-dom';
import Search from './Search';
import './App.css';

class App extends Component {
  render() {
    return (
      <div className="App">
        <p className="App-intro">
          Simple algolia search:
        </p>
        <InstantSearch appId="" apiKey="" indexName="my_index">
          <Search />
        </InstantSearch>

      </div>
    );
  }
}

export default App;
