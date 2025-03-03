# Configuration Options

> To see the available options, run gopsur init -h. The output will list all valid options, and their exact casing must be used (e.g., if an option is displayed as HTMX, it must be passed as HTMX).

## For init command

These will be asked as prompts when you run `init`. If an option is selected by a flag, it'll skip that prompt.

**Web Framework**
- Echo  
- Fiber
- Chi
```sh
# flag
--framework Echo
```

**Rendering Strategy**
- Templates  
- Seperate Client (eg. react,svelte,etc.)
```sh
# flag
--render Seperate
```

**Styling**
- Vanilla  
- Tailwind
```sh
# flag
--styling Tailwind
```

**UI Library** 
- Preline  
- DaisyUI
```sh
# flag
--ui DaisyUI
```

## Options via Flags

These options can only be enabled/used via flags.

**Extra Options**
- HTMX  
- Dockerfile
```sh
# flag
--extra Dockerfile
```