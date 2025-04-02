import { NavToggle } from "./modules/navToggle.js";
import { HeroHeader } from "./modules/loadHomePageHeroImage.js";
import { InitLightBox } from "./lightbox.js";
import { ObservationGroup } from "./modules/intersectionObserver.js";
import { ServiceCard } from "./modules/serviceCard.js";
import { Parallax } from "./modules/parallax.js";
import { BackToTop } from "./modules/backToTop.js";
import { formatTextarea } from "./modules/clearTextarea.js";
new NavToggle("navToggle");
new HeroHeader("homeHeader");
new ObservationGroup(".card.testimonial");
new ServiceCard("serviceList");
InitLightBox();
Parallax();
new formatTextarea("message");
new BackToTop("backToTop");
