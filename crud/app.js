const sendRequest = (request) =>
  fetch("/product", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(request),
  });

document.querySelector("#sendList").addEventListener("click", (e) => {
  e.preventDefault();
  const query = `
    query getProductList {
      list {
        id, name, price
      }
    }`;
  sendRequest({ query })
    .then((resp) => resp.json())
    .then((json) => {
      let data = JSON.stringify(json, null, "&#009;");
      data = data.replace(/\n/g, "<br>");
      document.querySelector("#productList").innerHTML = data;
    });
});

document.querySelector("#getProduct").addEventListener("submit", (e) => {
  e.preventDefault();
  const query = `
    query getProductList($id: Int!) {
      product(id: $id) {
        id,
        name,
        price,
        info
      }
    }`;
  sendRequest({
    query,
    variables: {
      id: document.querySelector("input[name=get-product-id]").value,
    },
  })
    .then((resp) => resp.json())
    .then((json) => {
      let data = JSON.stringify(json, null, "&#009;");
      data = data.replace(/\n/g, "<br>");
      document.querySelector("#productGet").innerHTML = data;
    });
});

document.querySelector("#createProduct").addEventListener("submit", (e) => {
  e.preventDefault();
  const query = `
    mutation createProduct($name: String!, $info: String, $price: Float!) {
      create(name: $name, info: $info, price: $price) {
        id,
        name,
        price,
        info
      }
    }`;
  sendRequest({
    query,
    variables: {
      name: document.querySelector("input[name=create-product-name]").value,
      info: document.querySelector("input[name=create-product-info]").value,
      price: document.querySelector("input[name=create-product-price]").value,
    },
  })
    .then((resp) => resp.json())
    .then((json) => {
      let data = JSON.stringify(json, null, "&#009;");
      data = data.replace(/\n/g, "<br>");
      document.querySelector("#productCreate").innerHTML = data;
    });
});

document.querySelector("#updateProduct").addEventListener("submit", (e) => {
  e.preventDefault();
  const query = `
    mutation updateProduct($id: Int!, $name: String!, $info: String, $price: Float!) {
      update(id: $id, name: $name, info: $info, price: $price) {
        id,
        name,
        price,
        info
      }
    }`;
  sendRequest({
    query,
    variables: {
      id: document.querySelector("input[name=update-product-id]").value,
      name: document.querySelector("input[name=update-product-name]").value,
      info: document.querySelector("input[name=update-product-info]").value,
      price: document.querySelector("input[name=update-product-price]").value,
    },
  })
    .then((resp) => resp.json())
    .then((json) => {
      let data = JSON.stringify(json, null, "&#009;");
      data = data.replace(/\n/g, "<br>");
      document.querySelector("#productUpdate").innerHTML = data;
    });
});

document.querySelector("#deleteProduct").addEventListener("submit", (e) => {
  e.preventDefault();
  const query = `
    mutation deleteProduct($id: Int!) {
      delete(id: $id) {
        id,
        name,
        price,
        info
      }
    }`;
  sendRequest({
    query,
    variables: {
      id: document.querySelector("input[name=delete-product-id]").value,
    },
  })
    .then((resp) => resp.json())
    .then((json) => {
      let data = JSON.stringify(json, null, "&#009;");
      data = data.replace(/\n/g, "<br>");
      document.querySelector("#productDelete").innerHTML = data;
    });
});
