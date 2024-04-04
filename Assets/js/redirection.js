function redirectionChronometree(url, delai) {
    setTimeout(function() {
      window.location.href = url;
    }, delai);
  }

  redirectionChronometree('/', 2000);