# MapParser

A Minecraft map item parser to turn map.dat files into pngs!

## Usage

Download and run mapparser.exe from [releases](/releases) inside the same directory as your map.dat file. Input your filename without the .dat extension and it will be converted to .png and output in the same directory.

You can also add this to your PATH enviroment variable so it can be run from the command line by just typing `mapparser`. Download and run the `addtopath.bat` in the same folder as your mapparser.exe and it will be copied to `C:\MapParser` and added to your PATH. Test it out by opening cmd and running `mapparser`.

## Contributing

PLEASE CONTRIBUTE!!! Copying the map color data alone from the Minecraft wiki page took 45 mins of straight copy paste, I would love it if other people helped to fill in the rest of the missing versions in [colormap.json](/blob/main/colormap.json) so we can have more version support. The Minecraft wiki page is [here](https://minecraft.wiki/w/Map_item_format). Also feel free to use [colormap.json](/blob/main/colormap.json) for any of your projects! I put it into a Json file so anyone can easily use it.
