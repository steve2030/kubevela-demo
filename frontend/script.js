// Fetch books from the backend API
const booksContainer = document.getElementById('books');
const loading = document.getElementById('loading');
const addBookBtn = document.getElementById('add-book-btn');
const modal = document.getElementById('book-modal');
const closeModal = document.querySelector('.close');
const bookForm = document.getElementById('book-form');
const modalTitle = document.getElementById('modal-title');

let currentBookId = null; // Track the book being edited

// Open modal for adding/editing books
addBookBtn.addEventListener('click', () => {
    currentBookId = null; // Reset for adding a new book
    modalTitle.textContent = 'Add New Book';
    bookForm.reset();
    modal.style.display = 'block';
});

// Close modal
closeModal.addEventListener('click', () => {
    modal.style.display = 'none';
});

// Handle form submission
bookForm.addEventListener('submit', (event) => {
    event.preventDefault();

    const bookData = {
        title: document.getElementById('title').value,
        author: document.getElementById('author').value,
        year: parseInt(document.getElementById('year').value),
        image: document.getElementById('image').value,
    };

    if (currentBookId) {
        // Edit existing book
        updateBook(currentBookId, bookData);
    } else {
        // Add new book
        addBook(bookData);
    }

    modal.style.display = 'none';
});

// Fetch and display books
function fetchBooks() {
    fetch('http://localhost:8080/api/books')
        .then(response => response.json())
        .then(data => {
            loading.style.display = 'none';
            renderBooks(data);
        })
        .catch(error => {
            loading.style.display = 'none';
            console.error('Error fetching books:', error);
            booksContainer.innerHTML = '<p style="color: red;">Failed to load books. Please try again later.</p>';
        });
}

// Render books
function renderBooks(books) {
    booksContainer.innerHTML = '';
    books.forEach(book => {
        const bookCard = document.createElement('div');
        bookCard.className = 'book-card';
        bookCard.innerHTML = `
            <img src="${book.image}" alt="${book.title}">
            <h3>${book.title}</h3>
            <p><strong>Author:</strong> ${book.author}</p>
            <p><strong>Year:</strong> ${book.year}</p>
            <div class="actions">
                <button class="edit-btn" data-id="${book.id}">Edit</button>
                <button class="delete-btn" data-id="${book.id}">Delete</button>
            </div>
        `;
        booksContainer.appendChild(bookCard);
    });

    // Add event listeners for edit and delete buttons
    document.querySelectorAll('.edit-btn').forEach(button => {
        button.addEventListener('click', (event) => {
            const bookId = event.target.getAttribute('data-id');
            editBook(bookId);
        });
    });

    document.querySelectorAll('.delete-btn').forEach(button => {
        button.addEventListener('click', (event) => {
            const bookId = event.target.getAttribute('data-id');
            deleteBook(bookId);
        });
    });
}

// Add a new book
function addBook(bookData) {
    fetch('http://localhost:8080/api/books', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(bookData),
    })
        .then(response => response.json())
        .then(() => {
            fetchBooks(); // Refresh the book list
        })
        .catch(error => {
            console.error('Error adding book:', error);
        });
}

// Edit a book
function editBook(bookId) {
    fetch(`http://localhost:8080/api/books/${bookId}`)
        .then(response => response.json())
        .then(book => {
            currentBookId = bookId;
            modalTitle.textContent = 'Edit Book';
            document.getElementById('title').value = book.title;
            document.getElementById('author').value = book.author;
            document.getElementById('year').value = book.year;
            document.getElementById('image').value = book.image;
            modal.style.display = 'block';
        })
        .catch(error => {
            console.error('Error fetching book:', error);
        });
}

// Update a book
function updateBook(bookId, bookData) {
    fetch(`http://localhost:8080/api/books/${bookId}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(bookData),
    })
        .then(response => response.json())
        .then(() => {
            fetchBooks(); // Refresh the book list
        })
        .catch(error => {
            console.error('Error updating book:', error);
        });
}

// Delete a book
function deleteBook(bookId) {
    if (confirm('Are you sure you want to delete this book?')) {
        fetch(`http://localhost:8080/api/books/${bookId}`, {
            method: 'DELETE',
        })
            .then(() => {
                fetchBooks(); // Refresh the book list
            })
            .catch(error => {
                console.error('Error deleting book:', error);
            });
    }
}

// Initial fetch of books
fetchBooks();