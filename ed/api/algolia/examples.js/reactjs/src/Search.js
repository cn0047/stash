import React from 'react';
import { Hits, SearchBox } from 'react-instantsearch-dom';
import Page from "./Page";

function Search() {
  return (
    <div className="container">
      <SearchBox />
      <br/>
      <Hits hitComponent={Page} />
    </div>
  );
}

export default Search;
