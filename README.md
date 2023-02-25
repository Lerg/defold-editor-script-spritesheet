# Spritesheet Editor Script for Defold

This editor script is intended to work with the Defold Exporter for TexturePacker https://github.com/Lerg/defold-texturepacker-exporter

The exporter generates an atlas, a spritesheet image and a JSON file with the spritesheet information.

This editor script takes the JSON file as an input and splits the spritesheet image into individual files for the Defold atlas.

To use this editor script:

- Install it as a dependency in your Defold project (don't forget to fetch).
- Click `Project -> Reload Editor Scripts`.
- Find the JSON file from the exporter.
- Right click the JSON file and select `Spritesheet: Split`

After a successful extraction, the spritesheet image and the JSON files are deleted.

---