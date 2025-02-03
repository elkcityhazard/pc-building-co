class HeroHeader {

    constructor(id="") {
        this.id = id
        this.homeHeroContainer = document.getElementById(this.id)
        this.events()
    }

    events() {
        document.addEventListener('load', this.loadImage())
    }

    loadImage() {
    if (typeof this.homeHeroContainer.dataset.bg == "undefined") {
        this.homeHeroContainer.classList.add('has-img')
        return
    }
    this.homeHeroContainer.style.background = `var(--sunset) url('${this.homeHeroContainer.dataset.bg}') no-repeat center / cover`
                
    }
}

export {
    HeroHeader
}