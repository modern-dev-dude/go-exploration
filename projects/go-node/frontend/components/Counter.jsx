import React from "react";

function Counter(props) {
  const [count, setCount] = React.useState(props.defaultNum);
  return (
    <div>
      <h2>{count}</h2>
      <div className="flex gap-4">
        <button
          type="button"
          className="rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
          onClick={() => setCount((state) => state + 1)}
        >
          Count up
        </button>

        <button
          type="button"
          className="rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
          onClick={() => setCount((state) => state - 1)}
        >
          Count down
        </button>
      </div>
    </div>
  );
}

export default Counter;
