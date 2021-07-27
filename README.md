# Participating media

Volumetric clouds rendering

Atmosphere scattering (Rayleigh / Mie)

## FEATURES

- RayTracing / RayMarching (light accumulation in clouds) ✓

- Noises (Perlin noise and Worley noise) for generating cloud densities ✓ 

- Atmospheric Scattering (Rayleigh and Mie) ✓ 

- Create animation (camera translation and camera rotation) ✓

- Manage multiple cloud (cloud superposition) ✓

- HeightMap (cloud height) ✓

- Multiple light sources ✓

- Parallelization using goroutine ✓

- SIMD instruction

## Results

![Alt Text](results/rotationsunsetfinal.gif)

![Alt Text](results/sunsettranslatefinal.gif)

![plot](results/daylight1.png)

![Alt Text](results/backward-anim-one-cloud.gif)

![plot](results/sunset1-5clouds.png)

![plot](results/30_clouds.png)

## Research Papers

Raymarch numerical integration:

https://arxiv.org/pdf/1609.05344.pdf


Participating media (good overview):

https://media.contentapi.ea.com/content/dam/eacom/frostbite/files/s2016-pbs-frostbite-sky-clouds-new.pdf


Raymarching Cloud (with sphere):

https://www.researchgate.net/publication/343404421\_The\_Current\_State\_of\_the\_Art\_in\_Real-Time\_Cloud\_Rendering\_With\_Raymarching


Rikard Olajos thesis (cloud rendering):

https://lup.lub.lu.se/luur/download?func=downloadFile&recordOId=8893256&fileOId=8893258


Dean Babić thesis (cloud rendering):

https://bib.irb.hr/datoteka/949019.Final\_0036470256\_56.pdf


Fredrik Haggstrom thesis (cloud rendering):

http://www.diva-portal.org/smash/get/diva2:1223894/FULLTEXT01.pdf


Rurik Hogfeldt thesis (cloud rendering):

https://publications.lib.chalmers.se/records/fulltext/241770/241770.pdf


Real Time cloud rendering:

http://www.markmark.net/PDFs/RTClouds\_HarrisEG2001.pdf


Multiple Clouds:

http://vterrain.org/Atmosphere/Clouds/


Texturing and Modeling (Gaz / Clouds):

http://elibrary.lt/resursai/Leidiniai/Litfund/Lithfund_leidiniai/IT/Texturing.and.Modeling.-.A.Procedural.Approach.3rd.edition.eBook-LRN.pdf


Transmittance explained:

https://pages.mtu.edu/~scarn/teaching/GE4250/transmission\_lecture.pdf


Atmospheric scattering (nvidia):

https://developer.nvidia.com/gpugems/gpugems2/part-ii-shading-lighting-and-shadows/chapter-16-accurate-atmospheric-scattering


Atmosphere Rendering (Sebastien Hillaire):

https://sebh.github.io/publications/egsr2020.pdf


Atmospheric Rendering:

https://core.ac.uk/download/pdf/55631247.pdf


Improved version of Perlin Noise (2002):

https://mrl.cs.nyu.edu/~perlin/paper445.pdf


Worley Noise:

https://weber.itn.liu.se/~stegu/TNM084-2017/worley-originalpaper.pdf


Higher order Ray marching:

https://graphics.unizar.es/papers/CGF2014\_higherorder.pdf


Cloud rendering final equation:

http://web.archive.org/web/20160604173317/http://freespace.virgin.net/hugo.elias/models/m\_clouds.htm


## Possible improvement (maybe Path Tracing):


Scattering in clouds:

https://hal.inria.fr/inria-00333007/document


Disney cloud rendering:

https://studios.disneyresearch.com/wp-content/uploads/2019/03/Deep-Scattering-paper.pdf


Pixar:

https://graphics.pixar.com/library/ProductionVolumeRendering/paper.pdf

## Commands


### Build

`./build.sh`


### Execute


Render Sunset:

`./volumetric-cloud sunset1`


Render Day:

`./volumetric-cloud daylight1`


Animation of the sunset:

`./volumetric-cloud sunsetanim`

`./generate_video.sh 24`


### Tests

`./tests.sh`


### View result (force aliasing)

`./view_img.sh <image-name>`
