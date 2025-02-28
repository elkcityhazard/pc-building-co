class NavToggle {
  constructor(navBtnId = "") {
    this.navBtnId = navBtnId;
    this.navBtn = document.getElementById(this.navBtnId);
    this.navBtn.setAttribute("aria-controls", "mainNav");
    this.mainNav = document.getElementById("mainNav") || null;
    this.navBarContainer = document.querySelector(".nav-bar");
    this.mainNavParents =
      this.mainNav || []
        ? this.mainNav.querySelectorAll("ul > li > a.parent")
        : [];
    this.navExpanded = false;
    this.currentWidth = window.innerWidth;
    this.hasActiveChild = false;
    this.events();
  }

  events() {
    this.mainNavParents.forEach(
      (parent) => (parent.nextElementSibling.style.height = 0 + "px"),
    );
    this.mainNavParents.forEach((parent) =>
      parent.addEventListener("click", this.handleParentClick.bind(this)),
    );
    this.navBtn.addEventListener("mouseenter", this.handleMouseIn.bind(this));
    this.navBtn.addEventListener("mouseleave", this.handleMouseOut.bind(this));
    this.navBtn.addEventListener("click", this.handleClick.bind(this));
    window.addEventListener(
      "resize",
      this.debounce(this.handleResize.bind(this)).bind(this),
    );
  }

  resetToDefault() {
    this.navBtn.classList.remove("active");
    this.mainNav.classList.remove("active");
    this.mainNav.style.height = 0 + "px";
    this.mainNav.setAttribute("aria-exanded", false);
    this.navBarContainer.classList.remove("active");
    this.navBarContainer.style.height = "auto";
    this.navExpanded = false;
    this.mainNavParents.forEach((parent) => {
      let subMenu = parent.closest("a").nextElementSibling;
      subMenu.classList.remove("active");
      subMenu.style.height = 0 + "px";
    });
  }

  debounce(func, timeout = 500) {
    let timer;
    return (...args) => {
      clearTimeout(timer);
      timer = setTimeout(() => {
        func.apply(this, args);
      }, timeout);
    };
  }

  toggleActiveClass(el) {
    el.classList.contains("active")
      ? el.classList.remove("active")
      : el.classList.add("active");

    el.setAttribute(
      "aria-expanded",
      el.classList.contains("active") ? true : false,
    );
  }

  handleParentClick(e) {
    e.preventDefault();
    this.hasActiveChild = !this.hasActiveChild;
    const subMenu = e.target.closest("a").nextElementSibling;
    // add active class to submenu
    this.toggleActiveClass(subMenu);
    // set submenu height
    subMenu.classList.contains("active")
      ? (subMenu.style.height = subMenu.scrollHeight + "px")
      : (subMenu.style.height = 0 + "px");
  }

  handleResize(e) {
    if (this.navExpanded && e.target.innerWidth > this.currentWidth) {
      this.resetToDefault();
    }
  }

  handleClick(e) {
    this.navExpanded = !this.navExpanded;

    let childHeight;
    if (this.mainNavParents.length) {
      this.mainNavParents.forEach((parent) => {
        let children = parent.nextElementSibling || [];
        if (!children.length) {
          childHeight = 0;
        }
        const { scrollHeight: currentChildHeight = 0 } = children;
        childHeight += currentChildHeight;
      });
    } else {
      this.mainNav.style.height = this.mainNav.scrollHeight + "px";
    }

    this.toggleActiveClass(this.navBarContainer);
    this.toggleActiveClass(this.mainNav);

    this.mainNav.style.height = this.mainNav.classList.contains("active")
      ? this.mainNav.scrollHeight + childHeight + "px"
      : 0 + "px";
    // close children
    if (this.hasActiveChild) {
      this.mainNavParents.forEach((parent) => parent.click());
    }
  }

  handleMouseIn(e) {
    this.paths = !this.paths
      ? e.target.getElementsByTagName("path")
      : this.paths;
    for (const path of this.paths) {
      path.classList.add(!path.classList.contains("active") ? "active" : null);
    }
  }

  handleMouseOut(e) {
    this.paths = !this.paths
      ? e.target.getElementsByTagName("path")
      : this.paths;
    for (const path of this.paths) {
      path.classList.remove(
        path.classList.contains("active") ? "active" : null,
      );
    }
  }
}

export { NavToggle };

