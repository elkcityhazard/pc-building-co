class BackToTop {
  constructor(id) {
    this.id = id;
    this.btn = document.getElementById(this.id) || null;

    this.events();
  }

  events() {
    if (!this.btn) {
      return null;
    }

    this.btn.addEventListener("click", this.handleClick.bind(this));
    document.addEventListener(
      "scroll",
      this.debounce(this.handleScroll.bind(this)).bind(this),
    );
    this.handleScroll();
  }

  handleClick(e) {
    window.scrollTo(0, 0);
  }

  handleScroll(e) {
    const { scrollTop = 0 } = document.documentElement;
    var body = document.body,
      html = document.documentElement;

    var height = Math.max(
      body.scrollHeight,
      body.offsetHeight,
      html.clientHeight,
      html.scrollHeight,
      html.offsetHeight,
    );
    if (scrollTop < height * 0.1) {
      this.btn.style.opacity = 0;
    }

    if (scrollTop > height * 0.1) {
      this.btn.style.opacity = 1;
    }

    this.btn.ontransitionend = function (e) {
      if (e.target.style.opacity == 0) {
        e.target.style.bottom = "-100%";
      } else {
        e.target.style.bottom = "1rem";
      }
    };
  }
  debounce(func, timeout = 200) {
    let timer;
    return (...args) => {
      clearTimeout(timer);
      timer = setTimeout(() => {
        func.apply(this, args);
      }, timeout);
    };
  }
}

export { BackToTop };
