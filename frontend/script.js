// Fetch books from the backend API
const booksContainer = document.getElementById('books');
const loading = document.getElementById('loading');

fetch('http://localhost:8080/api/books')
    .then(response => response.json())
    .then(data => {
        loading.style.display = 'none'; // Hide loading spinner

        // Render books
        data.forEach(book => {
            const bookCard = document.createElement('div');
            bookCard.className = 'book-card';
            bookCard.innerHTML = `
                <img src="${book.image}" alt="${book.title}">
                <h3>${book.title}</h3>
                <p><strong>Author:</strong> ${book.author}</p>
                <p><strong>Year:</strong> ${book.year}</p>
            `;
            booksContainer.appendChild(bookCard);
        });

        // Add search functionality
        const searchInput = document.getElementById('search');
        searchInput.addEventListener('input', (event) => {
            const searchTerm = event.target.value.toLowerCase();
            const books = document.querySelectorAll('.book-card');
            books.forEach(book => {
                const title = book.querySelector('h3').textContent.toLowerCase();
                const author = book.querySelector('p').textContent.toLowerCase();
                if (title.includes(searchTerm) || author.includes(searchTerm)) {
                    book.style.display = 'block';
                } else {
                    book.style.display = 'none';
                }
            });
        });
    })
    .catch(error => {
        loading.style.display = 'none'; // Hide loading spinner
        console.error('Error fetching books:', error);
        booksContainer.innerHTML = '<p style="color: red;">Failed to load books. Please try again later.</p>';
    });