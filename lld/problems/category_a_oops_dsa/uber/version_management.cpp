/*

The storage team is developing a version management system. Every night we built one or a few new releases. The versions could be compatible, which means the system could be upgraded without the data migration; otherwise, we have to upgrade the data before deploying the new release binaries (typically with a MapReduce job to rebuild the data). The compatibility is transitive, which means if version a is compatible with version b, version b is compatible with c, then a is compatible with c. If any of them are incompatible, it makes a to c incompatible. We already have a compatible check test (Jenkins job). After each release, we only run the compatible test against the previous version.

public void addNewVersion(int ver, boolean isCompatibleWithPrev)
{
}

```
public boolean isCompatible(int srcVer, int targetVer)
{
}

```

1 -> 2 -> 3 -> 4 -> 5 -> 6
T    T    F    T    T

```
    VersionCompatibilityManagement versions = new VersionCompatibilityManagement();
    versions.addNewVersion(1, false);
    versions.addNewVersion(2, true);
    versions.addNewVersion(3, true);
    versions.addNewVersion(4, false);
    versions.addNewVersion(5, true);
    versions.addNewVersion(6, true);
    assert(versions.isCompatible(1, 3) == true);
    assert(versions.isCompatible(3, 5) == false);
    assert(versions.isCompatible(4, 2) == false); // downgrade
    assert(versions.isCompatible(3, 3) == true); // upgrade to itself, always compatible

```

*/