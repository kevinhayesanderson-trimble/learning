Concrete Encoding of INformation:
-  physical representation of bits
    - should be small and inexpensive( we want lot of them)
    - stable, reliable, repeatable
    - ease of access and speed of manipulation
      - access, transform, combine, trnasmit and store
     
- representing bits with charged particles
  - voltage and current to encode information in wired communication
  - phase and frequency for wireless communication
 
Analog Signaling:
- example: in image conversion:
  - raster -> continous voltage signal or wave ->
  - sync signals to denote end of the scan signals
  - processing block:
    - input voltage, ouput voltage
    - using the pre-packaged functionality in the processing blocks
    - composistion of fucntions - tinker toy assembly
    - processing errors - introduced by real physical phenomenons(eg: manufacturing variations, electron spin)
   
Digital Abstractions:
- Key idea: encode only one bit of information, 2 values 0, 1
- uniform representation convention
- 

Combinational Device:
- Static Discipline, below four propertities:
  1,2 - input and output should be digital
  3 - functional specificaiton:
    - values of the outputs for every possible combination of valid inputs
  4 - timing specification:
    - upper bound of the time required to compute the output for set of valid inputs
- Combinational Digital System:
  - interconnected elements of combinational device
  - no directed cycles
 
Noise:
- noise margins voltage levels to indicate the valid input and output voltage representations
- this is to avoid the voltage in the forbidden zone
- noise immunity

Voltage transfer characteristics:
- plot of vout vs vin : static behaviour(not plotted against time)
- plot indicates - valid and invalid input and output voltages and forbidden zone
- GAIN > 1 and nonlinear:
  - gain is vout/vin Small input changes trigger massive output changes (amplification).
  - nonlinear is The circuit behaves differently at low, middle, and high input voltages, not a uniform line
  -  
- 
