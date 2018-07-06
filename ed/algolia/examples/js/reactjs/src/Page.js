import React from 'react';

function Page({ hit }) {
  return (
    <div className="hit-name">
      <span>{hit.text}</span>
    </div>
  );
}

export default Page;
