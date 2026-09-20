import { InitLightBox } from "./lightbox.js";
import { BackToTop } from "./modules/backToTop.js";
import { ObservationGroup } from "./modules/intersectionObserver.js";
import { HeroHeader } from "./modules/loadHomePageHeroImage.js";
import { NavToggle } from "./modules/navToggle.js";
import { Parallax } from "./modules/parallax.js";
import { ServiceCard } from "./modules/serviceCard.js";

new NavToggle("navToggle");
new HeroHeader("homeHeader");
new ObservationGroup(".card.testimonial");
new ServiceCard("serviceList");
InitLightBox();
Parallax();
new BackToTop("backToTop");
