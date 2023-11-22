console.log("hello from cart");

function getCart() {
    const cartJson = localStorage.getItem('cart') || '[]';
    return JSON.parse(cartJson);
}

function setCart(cart) {
    localStorage.setItem('cart', JSON.stringify(cart));
}

function initCart() {
    // Initialize an empty cart if it doesn't exist
    if (!getCart()) {
        setCart([]);
    }
}

function AddToCart(name,ProductId,price, tax, image,quantity) {
    console.log("product id is", ProductId);
    console.log("price is", price);
    console.log("ProductName is", name);
    console.log("tax is", tax);
    console.log("image is", image);
    const cart = getCart();
    // Check if the product is already in the cart
    const existingProduct = cart.find(item => item.id === ProductId);

    if (existingProduct) {
        existingProduct.quantity += quantity;
    } else {
        cart.push({ id: ProductId, name,price, quantity,tax,image });
    }
    setCart(cart);
    // Show the "Added to Cart" popup
    showAddedToCartPopup();
}




function UpdateCart(ProductId, quantity) {
    const cart = getCart();
    const existingProduct = cart.find(item => item.id === ProductId);

    if (existingProduct && quantity > 0) {
        existingProduct.quantity = quantity;
        setCart(cart);
    }
}

function RemoveFromCart(ProductId) {
    const cart = getCart();
    const index = cart.findIndex(item => item.id === ProductId);

    if (index !== -1) {
        cart.splice(index, 1);
        setCart(cart);
    }
}

function showAddedToCartPopup() {
    const addedToCartPopup = document.getElementById('added-to-cart-popup');
    addedToCartPopup.style.display = 'block';

    // Hide the popup after a few seconds (e.g., 3 seconds)
    setTimeout(function () {
        addedToCartPopup.style.display = 'none';
    }, 3000);
}


function GetCartTotal(cart){
    let total = 0
    cart.forEach(element => {
        total +=element.price*element.quantity;
    });
    return total
}

function GetCartTax(cart){
    let total = 0
    cart.forEach(element => {
        total +=element.tax * element.quantity;
    });
    return total
}


// Call initCart when the page loads to ensure an empty cart exists
initCart();
