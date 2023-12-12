import React, { useState, useEffect, useCallback } from "react";
import axios from "axios";

function AppMine(callback) {
  return (
    <>
      <button onClick={callback.callback}>Click to call callback</button>
    </>
  );
}
function App() {
  const [data, setData] = useState({ hits: [] });
  useEffect(() => {
    (async () => {
      const result = await axios(
        "https://hn.algolia.com/api/v1/search?query=redux"
      );
      setData(result.data);
    })();
  }, []);

  const funccc = useCallback(() => {
    console.log("data");
    console.log(data);
  });
  return (
    <>
      <AppMine callback={funccc} />
    </>

    // <ul>
    //   {data.hits.map((item) => (
    //     <li key={item.objectID}>
    //       <a href={item.url}>{item.title}</a>
    //     </li>
    //   ))}
    // </ul>
  );
}
export default App;
