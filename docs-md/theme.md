# Theme Support in Ellie

Ellie now supports dynamic theming for both CLI and GUI output. You can choose between **light**, **dark**, and **auto** (automatic detection) modes. The system also detects always-dark terminals (like Ghost) and applies the dark theme automatically unless overridden.

## Features
- **Light and Dark Modes:** Choose your preferred color scheme for better readability.
- **Auto Mode:** Automatically selects the best theme based on your terminal or environment.
- **Always-Dark Terminal Detection:** If you use a terminal that is always dark (e.g., Ghost), Ellie will use the dark theme by default.
- **Dynamic Styling:** All styled output (headers, errors, info, etc.) adapts to the selected theme.

## CLI Usage

### Show the Current Theme
```sh
ellie theme show
```

### Set the Theme
- **Dark mode:**
  ```sh
  ellie theme set dark
  ```
- **Light mode:**
  ```sh
  ellie theme set light
  ```
- **Auto (detect):**
  ```sh
  ellie theme set auto
  ```

### Invalid Theme Example
```sh
ellie theme set blue
# Output: Invalid theme. Use 'light', 'dark', or 'auto'.
```

## How It Works
- The theme affects all styled output in the CLI.
- When set to `auto`, Ellie checks your terminal environment and applies the best theme.
- You can change the theme at any time using the commands above.

## Code Usage (for Developers)
When writing Go code for Ellie, use the dynamic style getters for output:

```go
import "github.com/tacheraSasi/ellie/styles"

styles.GetHeaderStyle().Println("This is a header")
styles.GetErrorStyle().Println("This is an error message")
styles.GetSuccessStyle().Println("Success!")
```

This ensures your output always matches the user's selected theme.

## Backward Compatibility
The old style variables (e.g., `styles.HeaderStyle`) are still available for legacy code, but new code should use the dynamic getters for theme consistency.

---

For more information, see the main README or run `ellie theme show` in your terminal.
