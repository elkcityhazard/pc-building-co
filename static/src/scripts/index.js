import {NavToggle} from "./modules/navToggle.js"
import { HeroHeader } from "./modules/loadHomePageHeroImage.js";
import { InitLightBox } from "./lightbox.js";
import { ObservationGroup } from "./modules/intersectionObserver.js";
import { ServiceCard } from "./modules/serviceCard.js";

//window.addEventListener('load', function(e) {
    new NavToggle("navToggle");
    new HeroHeader('homeHeader')
    new ObservationGroup(".card.testimonial")
    new ServiceCard("serviceList")
    InitLightBox()
//})



