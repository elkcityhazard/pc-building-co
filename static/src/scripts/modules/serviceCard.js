class ServiceCard {
  constructor(id) {
    this.id = id;
    this.serviceList = document.getElementById(this.id) || null;
    if (!this.serviceList) return null;
    this.buttons = this.serviceList.querySelectorAll(
      "button.services__btn-btn",
    );
    this.cardContainer = this.serviceList.querySelector(
      ".service-listings__inner",
    );
    this.cards = this.cardContainer.querySelectorAll("service-listings-card");
    this.activeID = 0;
    this.events();
  }

  events() {
    if (!this.serviceList || !this.buttons) return null;
    this.buttons.forEach((btn) =>
      btn.addEventListener("click", this.handleBtnClick.bind(this)),
    );
  }

  async handleBtnClick(e) {
    const { id } = e.target.dataset;
    if (id == this.activeID) return;
    this.removeActive();
    e.target.classList.add("active");
    await new Promise((resolve) => {
      setTimeout(() => {
        resolve(id);
      }, 250);
    })
      .then((data) => this.addActiveToCard(data))
      .catch((err) => console.error(err.message));
    this.cardContainer.style.transform = "translateX(" + -100 * id + "%)";
    this.activeID = id;
  }

  removeActive() {
    this.buttons.forEach((button) => button.classList.remove("active"));
    this.cards.forEach((card) => {
      card.classList.remove("active");
    });
  }

  addActiveToCard(cardID) {
    this.cards.forEach((card) => {
      card.dataset.id === cardID ? card.classList.add("active") : null;
      card.dataset.id === cardID
        ? card.setAttribute("aria-expanded", true)
        : card.setAttribute("aria-expanded", false);
    });
  }
}

export { ServiceCard };

