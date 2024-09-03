# `definitions.yaml`

The [definitions.yaml](./definitions.yaml) file maps vehicle signals to the VSS schema and specifies how to convert and interpret these signals. For a detailed explanation of YAML files, you can refer to the [YAML documentation](https://yaml.org/spec/1.2/spec.html).

## vspecName

The `vspecName` field defines the VSS (Vehicle Signal Specification) name of the singal. The name must be definied in the CSV ouptut of the VSS specification. VSS is a standardized schema defined in [vspec](https://covesa.github.io/vehicle_signal_specification/). This schema is used for vehicle data, and DIMO has its own [fork of the VSS](https://github.com/DIMO-Network/VSS) definitions tailored to our specific needs. This fork includes additional fields and modifications relevant to DIMO's data model.

## conversions

The `conversions` section maps fields from the original data to the VSpec field. Each `conversion` entry specifies:

- **`originalName`**: The name of the field in the original data.
- **`originalType`**: The type of the field in the original data.
- **`isArray`**: Whether the field is an array or not.

Multiple `conversions` can be used if there are different field names or types that need to be mapped to a single VSpec field. When generating code, the system will search the document for each `originalName` until the first match is found. This allows flexibility in handling variations in field names or types from different data sources.

## requiredPrivileges

The `requiredPrivileges` field lists the privileges required to access the signal. This ensures that only users with the appropriate permissions can access sensitive or specific vehicle data.

## Example

Here's a breakdown of a sample entry in the `definitions.yaml` file:

```yaml
- vspecName: Vehicle.Chassis.Axle.Row1.Wheel.Left.Tire.Pressure
  conversions:
    - originalName: tires.frontLeft
      originalType: float64
      isArray: false
  requiredPrivileges:
    - VEHICLE_NON_LOCATION_DATA
```

- **`vspecName`**: This maps to the VSpec field for the left front tire pressure.
- **`conversions`**: The `originalName` is `tires.frontLeft` and the type is `float64`. This field is not an array.
- **`requiredPrivileges`**: Access requires the `VEHICLE_NON_LOCATION_DATA` privilege.
