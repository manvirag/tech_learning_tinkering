import data from "./data.json";

export function fetchData() {
  let catsPromise = fetchCats();
  return {
    cats: wrapPromise(catsPromise)
  };
}

// Suspense integrations like Relay implement
// a contract like this to integrate with React.
// Real implementations can be significantly more complex.
// Don't copy-paste this into your project!
function wrapPromise(promise) {
  let status = "pending";
  let result;
  let suspender = promise.then(
    (r) => {
      status = "success";
      result = r;
    },
    (e) => {
      status = "error";
      result = e;
    }
  );
  return {
    read() {
      if (status === "pending") {
        throw suspender;
      } else if (status === "error") {
        throw result;
      } else if (status === "success") {
        return result;
      }
    }
  };
}

function fetchCats() {
  console.log("fetch cats...");
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      console.log("fetched cats");
      resolve(data);
      // reject();
    }, 1100);
  });
}
