import express from "express";

const bookMap = {
  1: {
    id: 1,
    title: "Harry Potter",
    author: "J.K. Rowling",
    description: "This is a book about a young wizard",
  },
  2: {
    id: 2,
    title: "Game of Thrones",
    author: "George R. R. Martin",
    description: "This is a book about a mid-evil dictator",
  },
};

const app = express();
app.use(express.json({ limit: "2mb" }));

app.post("/api/book", (req, res) => {
  const id = req?.body?.bookId;
  if (!bookMap[id]) {
    res.status(404).json({ error: "Book not found" });
  }

  res.json(bookMap[id]);
});

app.listen(3000, () => {
  console.log("Server is running on port 3000");
});
