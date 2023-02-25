#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import json
import os
import sys
from PIL import Image

def split_spritesheet(project_root_path: str, spritesheet_json_path: str):
	is_success = False
	with open(spritesheet_json_path, 'r', encoding = 'utf-8') as spritesheet_json_file:
		spritesheet_json = json.load(spritesheet_json_file)
		spritesheet_json_spritesheetFilename = spritesheet_json['spritesheetFilename']
		spritesheet_json_sprites = spritesheet_json['sprites']
		if isinstance(spritesheet_json_spritesheetFilename, str) and isinstance(spritesheet_json_sprites, list):
			spritesheet_image_path = os.path.join(os.path.dirname(spritesheet_json_path), spritesheet_json_spritesheetFilename)
			with Image.open(spritesheet_image_path) as image:
				for sprite in spritesheet_json_sprites:
					sprite_image_path = os.path.join(project_root_path, *sprite['path'][1:].split('/'))
					box = (sprite['x'], sprite['y'], sprite['x'] + sprite['width'], sprite['y'] + sprite['height'])
					sprite = image.crop(box)
					os.makedirs(os.path.dirname(sprite_image_path), exist_ok = True)
					sprite.save(sprite_image_path, quality = 100)
					print('Extracted sprite:', sprite_image_path)
			os.remove(spritesheet_image_path)
			print('Deleted spritesheet:', spritesheet_image_path)
			is_success = True
	if is_success:
		os.remove(spritesheet_json_path)
		print('Deleted spritesheet JSON:', spritesheet_json_path)

if __name__ == '__main__':
	if len(sys.argv) == 2:
		split_spritesheet('.', sys.argv[1])
	else:
		print('Missing path to the sprite sheet JSON file as an argument.')
		print('Must be called from the root directory of the Defold project.')
