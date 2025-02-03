class ObservationGroup {
    constructor(className="") {
        this.className = className
        this.targets = document.querySelectorAll(this.className)
        this.galleryTargets = document.querySelectorAll('#galleryWork img')
        this.observer = new IntersectionObserver(this.observerCallback.bind(this), this.options)
        this.galleryObserver = new IntersectionObserver(this.galleryObserverCallback.bind(this), this.options)
        this.servicesImg = document.querySelectorAll(".services__img") || null
        this.servicesImgObserver = new IntersectionObserver(this.servicesImgCallback.bind(this), this.options)
        this.options = {
            root: document.documentElement,
            rootMargin: "0px",
            threshold: 1.0,
        }

        this.events()
    }


    events() {

        this.targets.forEach(target => {
            this.observer.observe(target)
        })

        this.galleryTargets.forEach(target => {
            this.galleryObserver.observe(target)
        })

        this.servicesImg.forEach(target => {
            this.servicesImgObserver.observe(target)
        })

    }

    servicesImgCallback(entries) {

        if (!this.servicesImg) {
            return null
        }
        entries.forEach(async (e,idx) => {
            if (e.isIntersecting) {

                await new Promise(resolve => {
                    setTimeout(() => resolve(),350)
                })
                
                    e.target.src = e.target.getAttribute('data-src')
                    if (e.target.complete) {
                     e.target.classList.add('fadeUp')
                    this.observer.unobserve(e.target)
                    } else {
                        e.target.addEventListener('load', (e) => {
                            e.target.classList.add('fadeUp')
                            e.target.removeEventListener('load', function(e) {
                                this.observer.unobserve(e.target)
                            })
                        })
                    }
                    
                
            }
        })
    }

    galleryObserverCallback(entries) {
        if (!this.galleryTargets) return null
        entries.forEach(async (e,idx) => {
            if (e.isIntersecting) {

                await new Promise(resolve => {
                    setTimeout(() => resolve(),350)
                })
                
                    e.target.src = e.target.getAttribute('data-src')
                    if (e.target.complete) {
                     e.target.classList.add('fadeUp')
                    this.observer.unobserve(e.target)
                    } else {
                        e.target.addEventListener('load', (e) => {
                            e.target.classList.add('fadeUp')
                            e.target.removeEventListener('load', function(e) {
                                this.observer.unobserve(e.target)
                            })
                        })
                    }
                    
                
            }
        })
    }

    observerCallback(entries) {
        if (!this.targets) return null
        entries.forEach((e,idx) => {
            if (e.isIntersecting) {
                setTimeout(() => {
                    e.target.classList.add('fadeUp')
                this.observer.unobserve(e.target)
                }, 333 *(1+idx))
                
            }
        })
    }
}


export {
    ObservationGroup
}