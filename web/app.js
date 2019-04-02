var baseURL = location.protocol + "//" + location.host + "/sites";

document.addEventListener("DOMContentLoaded", function() {
  loadSites();
  setInterval(function() {
    loadSites();
  }, 60000);
});

function isValidUrl(url) {
  return /^https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)$/.exec(
    url
  )
    ? true
    : false;
}

function showAlert(text) {
  if (document.querySelector("#alert")) {
    document.querySelector("#alert").classList.remove("hide");
    document.querySelector("#alert").innerHTML = text;
    setTimeout(function() {
      document.querySelector("#alert").classList.add("hide");
    }, 3000);
  }
}

function addSite(e) {
  url = document.querySelector("#urlText").value;
  if (!isValidUrl(url)) {
    showAlert("Please enter a valid url like http://www.google.com");
    return;
  }
  setTimeout(function() {
    document.querySelector("#alert");
  });
  fetch(baseURL, {
    method: "post",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ url: url })
  })
    .then(function(response) {
      loadSites();
    })
    .catch(err => {
      showAlert(err);
      console.log("Something went wrong", err);
    });
}
function loadSites() {
  fetch(baseURL)
    .then(function(response) {
      return response.json();
    })
    .then(function(data) {
      console.log(JSON.stringify(data));
      renderUrlStatus(data);
      if (document.querySelector("#last-checked-at")) {
        console.log(document.querySelector("#last-checked-at"));
        document.querySelector("#last-checked-at").innerHTML = new Date();
      }
    })
    .catch(err => {
      showAlert(err);
      console.log("Something went wrong", err);
    });
}
function removeSite(e) {
  const uuid = e.target.attributes["data-uuid"].value;
  fetch(baseURL + "/" + uuid, {
    method: "delete"
  })
    .then(function(response) {
      loadSites();
    })
    .catch(err => {
      showAlert(err);
      console.log("Something went wrong");
    });
}

function renderUrlStatus(data) {
  var urlListItems = "";
  data.forEach(function(site) {
    var statusIcon = site.isHealthy ? "&#9989;" : "&#10060;";
    urlListItems +=
      '<div class="container" data-uuid="' +
      site.uuid +
      '">' +
      '<div class="row">' +
      '<div class="col-md-6">' +
      site.url +
      "</div>" +
      '<div class="col-md-4">' +
      statusIcon +
      "</div>" +
      '<div class="col-md-2">' +
      '<a href="#" data-uuid=' +
      site.uuid +
      " onclick=removeSite(event)>Delete</a>" +
      "</div>" +
      "</div>" +
      "</div>";
  });
  console.log(urlListItems);
  document.querySelector("#url-list-container").innerHTML = urlListItems;
}
