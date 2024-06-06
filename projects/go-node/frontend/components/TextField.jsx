import React from "react";
export default function TextField() {
  const [text, setText] = React.useState("");
  return (
    <div>
      <label
        htmlFor="some-text"
        className="block text-sm font-medium leading-6 text-gray-900"
      >
        Some Text:
      </label>
      <div className="mt-2">
        <input
          type="text"
          name="some-text"
          id="some-text"
          onChange={(e) => setText(e.target.value)}
          value={text}
          className="block w-full rounded-md border-0 p-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
          placeholder="Placeholder"
          aria-describedby="some-text"
        />
      </div>
      <p className="mt-2 text-sm text-gray-500 p-2" id="some-text-description">
        {text.length === 0 ? "No text entered yet!" : text}
      </p>
    </div>
  );
}
