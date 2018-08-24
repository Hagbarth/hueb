(function() {
  function startLoading(button) {
    button.classList.add("is-loading");
  }

  function stopLoading(button) {
    button.classList.remove("is-loading");
  }

  function fetchHref(e) {
    e.preventDefault();

    startLoading(this);
    const href = this.getAttribute("href");
    fetch(href)
      .then(r => r.text())
      .then(r => {
        stopLoading(this);
        if (r !== "ok") {
          alert(r);
        }
      })
      .catch(err => {
        stopLoading(this);
        alert(err);
      });

    return false;
  }

  document
    .getElementById("turn-on-button")
    .addEventListener("click", fetchHref, false);

  document
    .getElementById("turn-off-button")
    .addEventListener("click", fetchHref, false);
})();
